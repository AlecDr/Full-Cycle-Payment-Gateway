import Link from "next/link";
import { LogOut } from "lucide-react";

interface HeaderProps {
  username?: string;
}

export function Header({ username = "usuário" }: HeaderProps) {
  return (
    <header className="w-full bg-white border-b border-gray-200">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <Link href="/dashboard" className="text-xl font-bold text-slate-800">
          Full Cycle Gateway
        </Link>
        <div className="flex items-center gap-4">
          <span className="text-slate-600">Olá, {username}</span>
          <Link
            href="/auth"
            className="flex items-center gap-2 bg-red-600 hover:bg-red-700 text-white px-3 py-1 rounded"
          >
            <LogOut size={16} />
            <span>Logout</span>
          </Link>
        </div>
      </div>
    </header>
  );
}
