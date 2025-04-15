import { Skeleton } from "@/components/ui/skeleton";

export default function PerfilLoading() {
  return (
    <div className="max-w-2xl mx-auto py-10 space-y-10">
      <div>
        <Skeleton className="h-8 w-40 mb-6" />

        <div className="space-y-4 border p-6 rounded-lg shadow-sm">
          <Skeleton className="h-6 w-48" />
          <div className="space-y-4">
            <div>
              <Skeleton className="h-4 w-24 mb-2" />
              <Skeleton className="h-10 w-full" />
            </div>
            <div>
              <Skeleton className="h-4 w-24 mb-2" />
              <Skeleton className="h-10 w-full" />
            </div>
            <div>
              <Skeleton className="h-4 w-24 mb-2" />
              <Skeleton className="h-24 w-full" />
            </div>
          </div>
        </div>
      </div>

      <div className="space-y-4 border p-6 rounded-lg shadow-sm">
        <Skeleton className="h-6 w-48" />
        <div className="space-y-4">
          <div>
            <Skeleton className="h-4 w-32 mb-2" />
            <Skeleton className="h-10 w-full" />
          </div>
          <div>
            <Skeleton className="h-4 w-40 mb-2" />
            <Skeleton className="h-10 w-full" />
          </div>
        </div>
      </div>

      <div className="flex flex-row justify-end">
        <Skeleton className="h-10 w-14 mr-4" />

        <Skeleton className="h-10 w-14" />
      </div>
    </div>
  );
}
