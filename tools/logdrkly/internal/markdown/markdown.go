package markdown

import (
	"io"
	"os"
	"strings"
)

type Writer struct {
	depth   int
	builder *strings.Builder
}

func NewWriter(depth int) *Writer {
	return &Writer{
		builder: &strings.Builder{},
		depth:   depth,
	}
}

func (w *Writer) WriteTableHeader(headers ...string) {
	w.WriteTableRow(headers...)
	w.Write("|-")
	for idx, header := range headers {
		w.Write(strings.Repeat("-", len(header)))

		if idx != len(headers)-1 {
			w.Write("-|-")
		}
	}
	w.Write("-|")
	w.WriteBlankLn()
}

func (w *Writer) WriteTableRow(items ...string) {
	w.Write("| ")
	w.Write(strings.Join(items, " | "))
	w.Write(" |")
	w.WriteBlankLn()
}

func (w *Writer) Write(text string) {
	// Builder never fails a Write. https://pkg.go.dev/strings#Builder.Write
	_, _ = w.builder.WriteString(text)
}

func (w *Writer) WriteLn(text string) {
	w.Write(text)
	w.WriteBlankLn()
}

func (w *Writer) WriteBlankLn() {
	w.Write("\n")
}

func (w *Writer) WriteSection(title string, content func()) {
	w.depth += 1
	w.Write(strings.Repeat("#", w.depth))
	w.Write(" ")
	w.WriteLn(title)
	w.WriteBlankLn()
	defer func() {
		w.depth -= 1
	}()

	content()
	w.WriteBlankLn()
}

func (w *Writer) AppendMarkdown(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(string(bytes))
	return nil
}

func (w *Writer) Save(outPath string) error {
	file, err := os.Create(outPath)
	if err != nil {
		return err
	}
	_, err = file.WriteString(w.builder.String())
	return err
}

func (w *Writer) String() string {
	return w.builder.String()
}
