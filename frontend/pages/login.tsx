import { useState } from "react";

export default function Page() {
  const [login, setLogin] = useState(true); // true = login, false = register
  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center">
      {login ? loginView() : registerView()}
    </div>
  );

  function loginView() {
    return (
      <div className="bg-zinc-300 w-screen md:w-1/3 sm:w-2/3 h-1/3 rounded-lg flex flex-col items-center text-black">
        <h1 className="text-2xl mb-4">Login</h1>
        {forms()}
        <div className="flex gap-2 mt-2 items-center">
          <FormButton id="login" value="Login" />
          <span
            className="hover:cursor-pointer"
            onClick={() => setLogin(!login)}
          >
            or register
          </span>
        </div>
      </div>
    );
  }
  function registerView() {
    return (
      <div className="bg-zinc-300 w-screen md:w-1/3 sm:w-2/3 h-1/3 rounded-lg flex flex-col items-center text-black">
        <h1 className="text-2xl mb-4">Register</h1>
        {forms()}
        <div className="flex gap-2 mt-2 items-center">
          <FormButton id="register" value="Register" />
          <span
            className="hover:cursor-pointer"
            onClick={() => setLogin(!login)}
          >
            or login
          </span>
        </div>
      </div>
    );
  }

  function forms() {
    return (
      <form
        onSubmit={(e) => e.preventDefault()}
        className="flex flex-col items-center"
      >
        <label htmlFor="email">Email</label>
        <FormInput type="email" name="email" id="email" />
        <label htmlFor="email">Password</label>
        <FormInput type="password" name="password" id="password" />
      </form>
    );
  }

  function FormInput({
    type,
    name,
    id,
  }: {
    type: string;
    name: string;
    id: string;
  }) {
    return (
      <input
        type={type}
        name={name}
        id={id}
        className="p-2 rounded-md bg-white"
      />
    );
  }
  function FormButton({ id, value }: { id: string; value: string }) {
    return (
      <input
        type="submit"
        id={id}
        value={value}
        className="p-2 rounded-md bg-white hover:cursor-pointer"
        onClick={() => {
          if (login) {
            console.log("test");
            // TODO: login logic
          } else {
            // TODO: register logic
          }
        }}
      />
    );
  }
}
