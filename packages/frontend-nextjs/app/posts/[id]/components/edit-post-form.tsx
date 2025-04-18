"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";
import { Post, updatePost } from "@/actions/api/posts/posts";
import RichTextEditor from "../../components/text-editor";
import { Button } from "@/components/ui/button";
import { useToast } from "@/hooks/use-toast";
import { Spinner } from "@/components/ui/spinner";

export default function EditPostForm({ initialPost }: { initialPost: Post }) {
  const [htmlContent, setHtmlContent] = useState(initialPost.content);
  const [isPrivate, setIsPrivate] = useState(!initialPost.is_public);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const { toast } = useToast();

  const router = useRouter();

  const handleSubmit = async () => {
    if (!htmlContent) return;
    setIsSubmitting(true);

    try {
      await updatePost(initialPost.id!, {
        content: htmlContent,
        is_public: !isPrivate,
      });
      toast({
        title: "Postagem atualizada com sucesso",
      });

      setIsSubmitting(false);

      const redirectRoute = isPrivate ? "/timeline/mine" : "/timeline";
      router.push(redirectRoute);
    } catch (error) {
      setIsSubmitting(false);
      console.error("Failed to update post:", error);
    }
  };

  return (
    <div className="flex flex-col">
      <RichTextEditor
        onChange={setHtmlContent}
        initialContent={initialPost.content}
      />
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
        className="mt-4 self-end w-[92px]"
        disabled={!htmlContent}
      >
        {isSubmitting ? <Spinner /> : "Atualizar"}
      </Button>
    </div>
  );
}
