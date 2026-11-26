import karyawans from "./karyawan.json";

function App() {
  return (
    <div className="">
      <div className="bg-blue-400 grid items-center justify-center pt-8">
        <h1 className="text-4xl flex items-center text-center text-white font-serif pb-6">
          Data of Employes
        </h1>
      </div>



<div className="flex items-center gap-10 pt-4 px-120 justify-self-end-safe col-auto">
          <button
          type="submit"
          className="bg-blue-500 hover:bg-blue-600 text-white p-3 rounded font-semibold  mt-4"
        >
          Add employe
        </button>
</div>



      <div className="card flex items-center justify-center gap-10 pt-10 px-4 ">
        {karyawans &&
          karyawans.map((karyawan) => {
            return (
              <div className="cotainer px-auto">
                <div
                  className=" p-4 justify-center col-auto rounded-lg bg-blue-200 shadow-lg"
                  key={karyawan.id}
                >
                  name: <strong>{karyawan.name}</strong> <br />
                  phone: {karyawan.phone} <br />
                  email: {karyawan.email} <br />
                  job: {karyawan.job}
                </div>
              </div>
            );
          })}
      </div>
    </div>
  );
}

export default App;
