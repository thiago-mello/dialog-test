"use client";

"use client";

import { useState } from "react";
import RichTextEditor from "../../components/text-editor";
import { Post, saveNewPost } from "@/actions/api/posts/posts";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Spinner } from "@/components/ui/spinner";

export default function NewPostForm() {
  const [htmlContent, setHtmlContent] = useState<string | undefined>();
  const [isPrivate, setIsPrivate] = useState<boolean>(false);
  const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
  const router = useRouter();

  /**
   * Handles the submission of a new post.
   *
   * This function performs the following steps:
   * 1. Validates that HTML content exists
   * 2. Sets submission state to loading
   * 3. Creates a new post object with content and privacy settings
   * 4. Saves the post via API call
   * 5. Resets submission state
   * 6. Redirects to appropriate timeline based on post privacy
   *
   * @returns {Promise<void>}
   */
  const handleSubmit = async () => {
    if (!htmlContent) {
      return;
    }
    setIsSubmitting(true);

    const post: Post = {
      content: htmlContent,
      is_public: !isPrivate,
    };

    await saveNewPost(post);
    setIsSubmitting(false);

    const redirectRoute = isPrivate ? "/timeline/mine" : "/timeline";
    router.replace(redirectRoute);
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
        className="mt-4 self-end w-[73px]"
        disabled={!htmlContent || isSubmitting}
      >
        {isSubmitting ? <Spinner /> : "Salvar"}
      </Button>
    </div>
  );
}
