//  <form
//         onSubmit={handleLogin}
//         className="relative z-20 flex flex-col bg-white p-10 rounded-lg shadow-xl w-full max-w-md gap-4"
//       >
//         <h1 className="text-2xl font-bold mb-2 text-center text-gray-900">Login to Quizit</h1>

//         {error && <p className="text-red-500 bg-red-100 border border-red-400 p-2 rounded text-center">{error}</p>}

//         <div>
//           <label className="block font-bold text-gray-700 mb-1">Username</label>
//           <input
//             type="text"
//             className="w-full border-2 border-gray-300 p-2 rounded focus:outline-none focus:border-blue-500"
//             value={username}
//             onChange={(e) => setUsername(e.target.value)}
//           />
//         </div>

//         <div className="mb-2">
//           <label className="block font-bold text-gray-700 mb-1">Password</label>
//           <input
//             type="password"
//             className="w-full border-2 border-gray-300 p-2 rounded focus:outline-none focus:border-blue-500"
//             value={password}
//             onChange={(e) => setPassword(e.target.value)}
//           />
//         </div>

//         <button
//           type="submit"
//           className="bg-blue-600 hover:bg-blue-700 text-white p-3 rounded font-bold transition mt-2"
//         >
//           Login
//         </button>
//       </form>