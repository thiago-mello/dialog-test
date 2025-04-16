import NewPostForm from "./components/post-form";

export default async function NewPostPage() {
  return (
    <div className="max-w-2xl mx-auto py-10">
      <h1 className="text-2xl font-bold mb-6">Nova Postagem</h1>
      <NewPostForm />
    </div>
  );
}
