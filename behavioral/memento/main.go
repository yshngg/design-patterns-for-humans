package main

import "fmt"

type EditorMemento struct {
	content string
}

func NewEditorMemento(content string) *EditorMemento {
	return &EditorMemento{
		content: content,
	}
}

func (m EditorMemento) Content() string {
	return m.content
}

type Editor struct {
	content string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) Type(words string) {
	e.content += words
}

func (e *Editor) Content() string {
	return e.content
}

func (e *Editor) Save() *EditorMemento {
	return NewEditorMemento(e.content)
}

func (e *Editor) Restore(memento *EditorMemento) {
	e.content = memento.Content()
}

func main() {
	editor := NewEditor()

	// 输入一些内容
	editor.Type("This is the first sentence.")
	editor.Type("This is second.")

	// 保存状态，之后可恢复到: This is the first sentence. This is second.
	saved := editor.Save()

	// 再输入一些内容
	editor.Type("And this is third.")

	// 输出: 保存之前的内容
	fmt.Println(editor.Content()) // This is the first sentence. This is second. And this is third.

	editor.Restore(saved)
	fmt.Println(editor.Content()) // This is the first sentence.This is second.
}
