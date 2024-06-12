import * as fsp from 'node:fs/promises';
import {
  LogCodes,
  Condition,
  Message,
  System
} from './ld_log_codes';

const reservedWords = [
  "break",
  "as",
  "any",
  "case",
  "implements",
  "boolean",
  "catch",
  "interface",
  "constructor",
  "class",
  "let",
  "declare",
  "const",
  "package",
  "get",
  "continue",
  "private",
  "module",
  "debugger",
  "protected",
  "require",
  "default",
  "public",
  "number",
  "delete",
  "static",
  "set",
  "do",
  "yield",
  "string",
  "else",
  "symbol",
  "enum",
  "type",
  "export",
  "from",
  "extends",
  "of",
  "false",
  "finally",
  "for",
  "function",
  "if",
  "import",
  "in",
  "instanceof",
  "new",
  "null",
  "return",
  "super",
  "switch",
  "this",
  "throw",
  "true",
  "try",
  "typeof",
  "var",
  "void",
  "while",
  "with",
];

function safeIdentifier(param: string): string {
  if (reservedWords.includes(param)) {
    return `_${param}`;
  }
  return param;
}

function safeMessage(message: Message): string {
  let updatedMessage = message.parameterized;
  for (let param of Object.keys(message.parameters ?? {})) {
    if (reservedWords.includes(param)) {
      updatedMessage = updatedMessage.replace(`\$\{${param}\}`, `\$\{_${param}\}`);
    }
  }
  return updatedMessage;
}

function capitalize(val: string) {
  return val.charAt(0).toUpperCase() + val.slice(1);
}

function makeObjectIdentifier(input: string): string {
  return safeIdentifier(capitalize(input));
}

function makeParams(message: Message): string {
  return Object.keys(message.parameters || {}).map(param => `${safeIdentifier(param)}: string`).join(', ');
}

async function main() {
  const definitions: LogCodes = JSON.parse(await fsp.readFile('../../definition/codes/codes.json', { encoding: 'utf8' }))

  const outputFile = await fsp.open('./src/codes.ts', 'w');
  let depth = 0;

  async function write(text: string) {
    await fsp.appendFile(outputFile, `${'  '.repeat(depth)}${text}`);
  }

  async function writeLn(text: string) {
    await write(`${text}\n`);
  }

  async function scoped(start: string, end: string, scope: () => Promise<void>): Promise<void> {
    await writeLn(start);
    depth++;
    await scope();
    depth--;
    await writeLn(end);
  }

  async function withDocComment(makeComments: () => Promise<void>) {
    await writeLn('/**');
    await makeComments();
    await writeLn('*/');
  }

  async function writeCommentLn(comment: string) {
    await writeLn(` * ${comment}`);
  }

  async function writeCommentBlank() {
    await writeLn(` *`);
  }

  async function writeCommentParamLn(name: string, def: string) {
    await writeCommentLn(`@param ${name} ${def}`);
  }

  async function writeMessageFunctionDocComment(definition: Condition) {
    await withDocComment(async () => {
      await writeCommentLn('Generate a log string for this code.');
      await writeCommentBlank();
      await writeCommentLn('This function will automatically include the log code.');

      for (let paramName of Object.keys(definition.message.parameters || {})) {
        const paramDef = definition.message.parameters![paramName];
        await writeCommentParamLn(safeIdentifier(paramName), paramDef);
      }
    });
  }

  async function writeCodeFunctionDocComment(definition: Condition) {
    await withDocComment(async () => {
      await writeCommentLn('Get the code for this condition.');
    });
  }

  async function writeGenericDocComment(definition: {description: string}) {
    await withDocComment(async () => {
      await writeCommentLn(definition.description);
    });
  }

  await writeLn('/* eslint-disable prettier/prettier */');
  await writeLn('/* eslint-disable @typescript-eslint/no-unused-vars */');
  await writeLn('/* eslint-disable @typescript-eslint/naming-convention */');
  await writeLn('// This code is automatically generated and should not be manually edited.');
  await writeLn('');

  await scoped('export const Logs = {', '}', async () => {
    for (let [systemName, system] of Object.entries(definitions.systems)) {
      await writeGenericDocComment(system);
      await scoped(`${makeObjectIdentifier(systemName)}: {`, `},`, async () => {
        for (let [clsName, cls] of Object.entries(definitions.classes)) {
          await writeGenericDocComment(cls);
          await scoped(`${capitalize(clsName)}: {`, '},', async () => {
            for (let [conditionCode, condition] of Object.entries(definitions.conditions)) {
              if (condition.system == system.specifier && condition.class == cls.specifier) {
                await writeCondition(condition, conditionCode);
              }
            }
          });
        }
      });
    }
  });

  async function writeCondition(condition: Condition, conditionCode: string) {
    await writeGenericDocComment(condition);
    await scoped(`${makeObjectIdentifier(condition.name)}: {`, '},', async () => {
      await writeMessageFunctionDocComment(condition);
      await scoped(`message:(${makeParams(condition.message)}) => {`, '},', async () => {
        await writeLn(`return \`${conditionCode} ${safeMessage(condition.message)}\`;`);
      });

      await writeCodeFunctionDocComment(condition);
      await scoped(`code:() => {`, '},', async () => {
        await writeLn(`return '${conditionCode}';`);
      });
    });
  }
}

main();
