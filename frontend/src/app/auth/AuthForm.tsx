import { redirect } from "next/navigation";
import { ArrowRight, Info } from "lucide-react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { cookies } from "next/headers";

export async function loginAction(formData: FormData) {
  "use server";

  const apiKey = formData.get("apiKey");

  const cookiesStore = await cookies();
  cookiesStore.set("apiKey", apiKey as string);

  redirect("/invoices");
}

export function AuthForm() {
  return (
    <form action={loginAction} className="space-y-4">
      <div className="space-y-2">
        <label htmlFor="apiKey" className="text-sm font-medium text-slate-700">
          API Key
        </label>
        <div className="relative">
          <Input
            id="apiKey"
            type="text"
            name="apiKey"
            placeholder="Digite sua API Key"
            required
          />
          <Button
            type="submit"
            size="icon"
            className="absolute right-0 top-0 rounded-l-none"
          >
            <ArrowRight size={18} />
          </Button>
        </div>
      </div>

      <div className="bg-blue-50 border border-blue-200 rounded-md p-4 flex gap-3">
        <Info size={20} className="text-blue-600 shrink-0 mt-0.5" />
        <div className="space-y-1">
          <h4 className="font-medium text-blue-700">Como obter uma API Key?</h4>
          <p className="text-sm text-slate-700">
            Para obter sua API Key, você precisa criar uma conta de comerciante.
            Entre em contato com nosso suporte para mais informações.
          </p>
        </div>
      </div>
    </form>
  );
}
