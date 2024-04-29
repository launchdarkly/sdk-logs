import * as fsp from 'node:fs/promises';
import {
  LogCodes,
  Condition,
  Message
} from './ld_log_codes';

function makeParams(message: Message): string {
  return Object.keys(message.parameters || {}).map(param => `${param}: string`).join(', ');
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
      await writeCommentLn(definition.description);
      await writeCommentBlank();
      await writeCommentLn('Generate a log string for this code.');
      await writeCommentBlank();
      await writeCommentLn('This function will automatically include the log code.');

      for (let paramName of Object.keys(definition.message.parameters || {})) {
        const paramDef = definition.message.parameters![paramName];
        await writeCommentParamLn(paramName, paramDef);
      }
    });
  }

  async function writeCodeFunctionDocComment(definition: Condition) {
    await withDocComment(async () => {
      await writeCommentLn(definition.description);
      await writeCommentBlank();
      await writeCommentLn('Get the code for this condition.');
    });
  }

  await writeLn('/* eslint-disable prettier/prettier */');
  await writeLn('/* eslint-disable @typescript-eslint/no-unused-vars */');
  await writeLn('/* eslint-disable @typescript-eslint/naming-convention */');
  await writeLn('// This code is automatically generated and should not be manually edited.');
  await writeLn('');

  for (let [conditionCode, condition] of Object.entries(definitions.conditions)) {
    const [systemName, sys] = Object.entries(definitions.systems).find(([_, value]) => {
      return value.specifier == condition.system
    })!;

    const [className, cls] = Object.entries(definitions.classes).find(([_, value]) => {
      return value.specifier == condition.class
    })!;

    await writeCondition(condition, systemName, className, conditionCode);
  }


  async function writeCondition(condition: Condition, systemName: string, className: string, conditionCode: string) {
    await writeMessageFunctionDocComment(condition);
    await scoped(`export function ${systemName}_${className}_${condition.name}(${makeParams(condition.message)}): string {`, '}', async () => {
      await writeLn(`return \`${conditionCode} ${condition.message.parameterized}\`;`);
    });

    await writeCodeFunctionDocComment(condition);
    await scoped(`export function ${systemName}_${className}_${condition.name}_code(): string {`, '}', async () => {
      await writeLn(`return '${conditionCode}';`);
    });
  }
}

main();
