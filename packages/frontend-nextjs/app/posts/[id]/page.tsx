import { getPost } from "@/actions/api/posts/posts";
import EditPostForm from "./components/edit-post-form";
import { notFound } from "next/navigation";

export default async function EditPostPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const id = (await params).id;
  const post = await getPost(id);
  if (!post) {
    notFound();
  }

  return (
    <div className="max-w-2xl mx-auto py-10">
      <h1 className="text-2xl font-bold mb-6">Editar Postagem</h1>
      <EditPostForm initialPost={post} />
    </div>
  );
}
