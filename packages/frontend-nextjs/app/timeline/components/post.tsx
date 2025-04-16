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
import { Edit, Heart, Trash2 } from "lucide-react";
import { deletePost, PostProjection } from "@/actions/api/posts/posts";
import { likePost } from "@/actions/api/posts/likes";
import Link from "next/link";
import {
  AlertDialog,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
  AlertDialogAction,
  AlertDialogDescription,
} from "@/components/ui/alert-dialog";

async function toggleLike(postId: string, isLiked: boolean): Promise<void> {
  await likePost(postId, !isLiked);
}

interface PostItemProps {
  post: PostProjection;
  currentUserId?: string;
  onPostDeleted?: (postId: string) => void;
}

export function PostItem({
  post: initialPost,
  currentUserId,
  onPostDeleted,
}: PostItemProps) {
  const [post, setPost] = useState(initialPost);
  const [isLiking, setIsLiking] = useState(false);
  const [isDeleting, setIsDeleting] = useState(false);

  const isOwnPost = currentUserId && post.user.id === currentUserId;
  const postWasUpdated = post.created_at !== post.updated_at;

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

  const handleDelete = async () => {
    setIsDeleting(true);
    try {
      await deletePost(post.id);
      if (onPostDeleted) {
        onPostDeleted(post.id);
      }
    } catch (error) {
      console.error("Erro ao excluir o post:", error);
    } finally {
      setIsDeleting(false);
    }
  };

  return (
    <Card className="mb-4">
      <CardContent className="pt-6">
        <div className="flex justify-between items-start mb-1">
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

          <div className="flex items-center gap-2">
            {isOwnPost && (
              <>
                <Link href={`/posts/${post.id}`} passHref>
                  <Button variant="ghost" size="icon" className="h-8 w-8">
                    <Edit size={16} className="text-gray-500" />
                    <span className="sr-only">Editar</span>
                  </Button>
                </Link>

                <AlertDialog>
                  <AlertDialogTrigger asChild>
                    <Button variant="ghost" size="icon" className="h-8 w-8">
                      <Trash2 size={16} className="text-gray-500" />
                      <span className="sr-only">Excluir</span>
                    </Button>
                  </AlertDialogTrigger>
                  <AlertDialogContent>
                    <AlertDialogHeader>
                      <AlertDialogTitle>Excluir post</AlertDialogTitle>
                      <AlertDialogDescription>
                        Tem certeza que deseja excluir este post? Esta ação não
                        pode ser desfeita.
                      </AlertDialogDescription>
                    </AlertDialogHeader>
                    <AlertDialogFooter>
                      <AlertDialogCancel>Cancelar</AlertDialogCancel>
                      <AlertDialogAction
                        onClick={handleDelete}
                        disabled={isDeleting}
                      >
                        {isDeleting ? "Excluindo..." : "Excluir"}
                      </AlertDialogAction>
                    </AlertDialogFooter>
                  </AlertDialogContent>
                </AlertDialog>
              </>
            )}
          </div>
        </div>

        <div className="text-xs text-gray-500 mb-3">
          {`${postWasUpdated ? "Atualizado" : "Criado"} ${formattedDate}`}
        </div>

        <div
          className="prose prose-sm max-w-none post-content"
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
