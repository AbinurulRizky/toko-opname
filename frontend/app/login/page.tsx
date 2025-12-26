"use client"
import { useState } from "react"
import { LoginResponse, ApiError } from "@/types/auth";

export default function Login() {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");

  const handleLogin = async() => {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/login`, {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ email, password }),
      });

      const data = await response.json();

      if(response.ok) {
        const successData = data as LoginResponse;
        console.log(successData.token)
        console.log(successData.role)

        localStorage.setItem("token", successData.token)
        alert(successData.message)
      } else {
        const errorData = data as ApiError
        alert(errorData.error)
      }
    } catch (error) {
      console.error("Gagal koneksi ke server", error)
    }
  };


  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-600 via-indigo-600 to-purple-700 px-4">
      <div className="w-full max-w-md backdrop-blur-lg bg-white/20 border border-white/30 rounded-2xl shadow-2xl p-8 text-white">

        <h1 className="text-3xl font-bold text-center mb-2">
          Welcome
        </h1>
        <p className="text-center text-white/80 mb-6">
          Login untuk melanjutkan
        </p>

        <div className="mb-4">
          <label className="block text-sm mb-1">Email</label>
          <input
            type="email"
            placeholder="Masukkan email"
            value={email}
            className="w-full px-4 py-2 rounded-lg bg-white/90 text-gray-800 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            onChange={(e) => setEmail(e.target.value)}

          />
        </div>

        <div className="mb-6">
          <label className="block text-sm mb-1">Password</label>
          <input
            type="password"
            value={password}
            placeholder="Masukkan password"
            className="w-full px-4 py-2 rounded-lg bg-white/90 text-gray-800 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            onChange={(e) => setPassword(e.target.value)}
          />
        </div>

        <button
          className="w-full mt-6 py-3 rounded-xl bg-white text-indigo-700 font-bold hover:bg-gray-100 hover:cursor-pointer transition"
          onClick={handleLogin}
        >
          Login
        </button>
      </div>
    </div>
  );
}