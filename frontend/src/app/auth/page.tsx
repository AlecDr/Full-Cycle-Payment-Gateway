import type React from "react";

import {
  Card,
  CardContent,
  CardHeader,
  CardTitle,
  CardDescription,
} from "@/components/ui/card";
import { AuthForm } from "./AuthForm";

export default function AuthPage() {
  return (
    <div className="min-h-screen flex flex-col bg-gray-50">
      <header className="w-full bg-white border-b border-gray-200">
        <div className="container mx-auto px-4 py-4">
          <h1 className="text-xl font-bold text-slate-800">
            Full Cycle Gateway
          </h1>
        </div>
      </header>

      <main className="flex-1 flex items-center justify-center p-4">
        <Card className="w-full max-w-md">
          <CardHeader>
            <CardTitle className="text-center">Autenticação Gateway</CardTitle>
            <CardDescription className="text-center">
              Insira sua API Key para acessar o sistema
            </CardDescription>
          </CardHeader>
          <CardContent>
            <AuthForm />
          </CardContent>
        </Card>
      </main>
    </div>
  );
}
