"use client";

"use client";

import { useState } from "react";
import RichTextEditor from "../../components/text-editor";
import { Post, saveNewPost } from "@/actions/api/posts/posts";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";

export default function NewPostForm() {
  const [htmlContent, setHtmlContent] = useState<string | undefined>();
  const [isPrivate, setIsPrivate] = useState<boolean>(false);
  const router = useRouter();

  const handleSubmit = async () => {
    if (!htmlContent) {
      return;
    }

    const post: Post = {
      content: htmlContent,
      is_public: !isPrivate,
    };
    const response = await saveNewPost(post);
    router.replace(`/posts/${response.id}`);
  };

  return (
    <div className="flex flex-col">
      <RichTextEditor onChange={setHtmlContent} />

      <div className="mt-4">
        <label>
          <input
            type="checkbox"
            checked={isPrivate}
            onChange={(e) => setIsPrivate(e.target.checked)}
            className="mr-1"
          />
          Tornar postagem privada
        </label>
      </div>

      <Button
        onClick={handleSubmit}
        className="mt-4 self-end"
        disabled={!htmlContent}
      >
        Enviar
      </Button>
    </div>
  );
}
