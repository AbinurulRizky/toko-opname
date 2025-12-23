export default function Login() {
   return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-600 via-indigo-600 to-purple-700 px-4">
      <div className="w-full max-w-md backdrop-blur-lg bg-white/20 border border-white/30 rounded-2xl shadow-2xl p-8 text-white">
        
        <h1 className="text-3xl font-bold text-center mb-2">
          Welcome 👋
        </h1>
        <p className="text-center text-white/80 mb-6">
          Login untuk melanjutkan
        </p>

        {/* Username */}
        <div className="mb-4">
          <label className="block text-sm mb-1">Username</label>
          <input
            type="text"
            placeholder="Masukkan username"
            className="w-full px-4 py-2 rounded-lg bg-white/90 text-gray-800 focus:outline-none focus:ring-2 focus:ring-indigo-500"
           
          />
        </div>

        {/* Password */}
        <div className="mb-6">
          <label className="block text-sm mb-1">Password</label>
          <input
            type="password"
            placeholder="Masukkan password"
            className="w-full px-4 py-2 rounded-lg bg-white/90 text-gray-800 focus:outline-none focus:ring-2 focus:ring-indigo-500"
            
          />
        </div>

      <button
  type="button"
  className="w-full mt-6 py-3 rounded-xl bg-white text-indigo-700 font-bold hover:bg-gray-100 transition"
>
  Login
</button>

  




       
          <p className="text-center mt-4 text-sm text-white">
          
          </p>
       
      </div>
    </div>
  );
}