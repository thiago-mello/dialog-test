"use client";

import { useState } from "react";
import { formatDistanceToNow } from "date-fns";
import { ptBR } from "date-fns/locale";
import {
  HoverCard,
  HoverCardContent,
  HoverCardTrigger,
} from "@/components/ui/hover-card";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Heart } from "lucide-react";
import { PostProjection } from "@/actions/api/posts/posts";
import { likePost } from "@/actions/api/posts/likes";

async function toggleLike(postId: string, isLiked: boolean): Promise<void> {
  await likePost(postId, !isLiked);
}

export function PostItem({ post: initialPost }: { post: PostProjection }) {
  const [post, setPost] = useState(initialPost);
  const [isLiking, setIsLiking] = useState(false);

  const formattedDate = formatDistanceToNow(new Date(post.updated_at), {
    addSuffix: true,
    locale: ptBR,
  });

  const handleLikeClick = async () => {
    if (isLiking) return;

    setIsLiking(true);

    // Optimistically update UI
    setPost((prevPost) => ({
      ...prevPost,
      like_count: prevPost.user_liked_this_post
        ? prevPost.like_count - 1
        : prevPost.like_count + 1,
      user_liked_this_post: !prevPost.user_liked_this_post,
    }));

    try {
      await toggleLike(post.id, post.user_liked_this_post);
    } catch (error) {
      // Revert changes on error
      setPost(initialPost);
    } finally {
      setIsLiking(false);
    }
  };

  return (
    <Card className="mb-4">
      <CardContent className="pt-6 post-content">
        <div className="flex justify-between items-start mb-3">
          <div className="flex items-center">
            <HoverCard>
              <HoverCardTrigger asChild>
                <span className="font-medium text-lg hover:underline cursor-pointer">
                  {post.user.name}
                </span>
              </HoverCardTrigger>
              <HoverCardContent className="w-64">
                <div className="space-y-2">
                  <h4 className="font-semibold">{post.user.name}</h4>
                  <p className="text-sm text-gray-500">{post.user.bio}</p>
                </div>
              </HoverCardContent>
            </HoverCard>
          </div>
          <span className="text-xs text-gray-500">
            Atualizado {formattedDate}
          </span>
        </div>

        <div
          className="prose prose-sm max-w-none"
          dangerouslySetInnerHTML={{ __html: post.content }}
        />
      </CardContent>

      <CardFooter className="pt-0">
        <div className="flex items-center">
          <Button
            variant="ghost"
            size="sm"
            className="flex items-center gap-1 text-gray-600"
            onClick={handleLikeClick}
            disabled={isLiking}
          >
            <Heart
              size={18}
              className={
                post.user_liked_this_post ? "fill-red-500 text-red-500" : ""
              }
            />
            <span className="ml-1">
              {post.like_count} {post.like_count === 1 ? "curtida" : "curtidas"}
            </span>
          </Button>
        </div>
      </CardFooter>
    </Card>
  );
}
