"use client";

import { useEditor, EditorContent, BubbleMenu } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Link from "@tiptap/extension-link";
import Image from "@tiptap/extension-image";
import CharacterCount from "@tiptap/extension-character-count";
import {
  Bold,
  Italic,
  Strikethrough,
  Code,
  List,
  ListOrdered,
  Image as ImageIcon,
  Link2,
} from "lucide-react";
import { Level } from "@tiptap/extension-heading";

interface EditorProps {
  onChange: (value: string) => void;
  initialContent?: string;
}

const RichTextEditor = ({ onChange, initialContent }: EditorProps) => {
  const editor = useEditor({
    extensions: [
      StarterKit.configure({
        codeBlock: false,
      }),
      Link.configure({
        openOnClick: false,
      }),
      Image.configure({
        inline: true,
        allowBase64: false,
      }),
      CharacterCount.configure({
        limit: 6000,
      }),
    ],
    content: initialContent ?? "",
    onUpdate: ({ editor }) => {
      onChange(editor.getHTML());
    },
  });

  if (!editor) {
    return null;
  }

  const addImage = () => {
    const url = window.prompt("URL da imagem");
    if (url) {
      editor.chain().focus().setImage({ src: url }).run();
    }
  };

  const setLink = () => {
    const url = window.prompt("URL do link");
    if (url) {
      editor.chain().focus().toggleLink({ href: url }).run();
    } else if (url === null) {
      editor.chain().focus().unsetLink().run();
    }
  };

  return (
    <div className="border rounded-lg bg-white">
      {/* Menu Bar */}
      <div className="flex flex-wrap gap-2 p-2 border-b">
        {/* Paragraph and Headings */}
        <select
          value={editor.getAttributes("heading").level}
          onChange={(e) => {
            const level = parseInt(e.target.value);
            if (level > 0 && level <= 6) {
              editor
                .chain()
                .focus()
                .toggleHeading({ level: level as Level })
                .run();
            } else {
              editor.chain().focus().setParagraph().run();
            }
          }}
          className="p-1 rounded border"
        >
          <option value="0">Corpo do Texto</option>
          <option value="1">Cabeçalho 1</option>
          <option value="2">Cabeçalho 2</option>
          <option value="3">Cabeçalho 3</option>
        </select>

        {/* Basic formatting */}
        <button
          onClick={() => editor.chain().focus().toggleBold().run()}
          className={`p-1 rounded ${
            editor.isActive("bold") ? "bg-gray-200" : ""
          }`}
        >
          <Bold className="w-5 h-5" />
        </button>
        <button
          onClick={() => editor.chain().focus().toggleItalic().run()}
          className={`p-1 rounded ${
            editor.isActive("italic") ? "bg-gray-200" : ""
          }`}
        >
          <Italic className="w-5 h-5" />
        </button>
        <button
          onClick={() => editor.chain().focus().toggleStrike().run()}
          className={`p-1 rounded ${
            editor.isActive("strike") ? "bg-gray-200" : ""
          }`}
        >
          <Strikethrough className="w-5 h-5" />
        </button>

        {/* Lists */}
        <button
          onClick={() => editor.chain().focus().toggleBulletList().run()}
          className={`p-1 rounded ${
            editor.isActive("bulletList") ? "bg-gray-200" : ""
          }`}
        >
          <List className="w-5 h-5" />
        </button>
        <button
          onClick={() => editor.chain().focus().toggleOrderedList().run()}
          className={`p-1 rounded ${
            editor.isActive("orderedList") ? "bg-gray-200" : ""
          }`}
        >
          <ListOrdered className="w-5 h-5" />
        </button>

        {/* Code */}
        <button
          onClick={() => editor.chain().focus().toggleCode().run()}
          className={`p-1 rounded ${
            editor.isActive("code") ? "bg-gray-200" : ""
          }`}
        >
          <Code className="w-5 h-5" />
        </button>

        {/* Link and Image */}
        <button
          onClick={setLink}
          className={`p-1 rounded ${
            editor.isActive("link") ? "bg-gray-200" : ""
          }`}
        >
          <Link2 className="w-5 h-5" />
        </button>
        <button onClick={addImage} className="p-1 rounded">
          <ImageIcon className="w-5 h-5" />
        </button>
      </div>

      {/* Editor Content */}
      <div className="p-4 min-h-[300px] flex">
        <EditorContent
          editor={editor}
          aria-expanded={true}
          className="flex flex-col flex-1 post-content"
        />
      </div>

      {/* Bubble Menu for Links */}
      {editor && (
        <BubbleMenu
          editor={editor}
          tippyOptions={{ duration: 100 }}
          className="bg-white p-1 rounded shadow-md border flex gap-1"
        >
          <button
            onClick={setLink}
            className={`p-1 rounded ${
              editor.isActive("link") ? "bg-gray-200" : ""
            }`}
          >
            <Link2 className="w-5 h-5" />
          </button>
        </BubbleMenu>
      )}
    </div>
  );
};

export default RichTextEditor;
