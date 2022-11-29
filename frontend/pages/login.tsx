import {useState} from "react";
import Image from 'next/image'
import {Inter} from '@next/font/google'
import Head from "next/head";

import {ToastContainer, toast} from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const inter = Inter({
  weight: ['200', '400'],
  subsets: ['latin'],
})

export default function Page() {
  const [login, setLogin] = useState(true); // true = login, false = register
  return (
    <div>
      <Head>
        <title>Kioku</title>
        <meta name="description" content="Kioku"/>
        <link rel="icon" href="/favicon.ico"/>
      </Head>
      <div className="w-screen h-screen select-none flex items-center">
        <div
          className="w-full h-fit flex flex-col items-center justify-center md:justify-evenly md:flex-row rounded-3xl bg-blue-50 m-10 min-w-max">
          <div className="flex flex-col items-center mb-0 md:mb-10 md:mr-2 md:mr-2 m-10 rounded-lg">
            <Image src="/kioku-logo.svg" alt="Kioku Logo" width={320} height={180}/>
            <span
              className={`${inter.className} font-extralight mt-5 tracking-[0.5em] text-6xl indent-[0.5em]`}>kioku</span>
          </div>
          {formView()}
        </div>
      </div>
      <ToastContainer position="bottom-center" autoClose={3000} hideProgressBar pauseOnFocusLoss/>
    </div>
  );

  function formView() {
    return (
      <div
        className={`bg-[#9EADC8] w-2/3 md:w-1/3 h-fit rounded-3xl flex flex-col items-center text-black ${inter.className} min-w-max mr-10 md:mr-2 m-10`}>
        <h1 className="text-2xl mt-5 mb-4">{login ? "Login" : "Register"}</h1>
        {forms()}
      </div>
    );
  }

  function loginButton() {
    return (
      <>
        <FormButton id="login" value="Login"/>
        <span className="hover:cursor-pointer" onClick={() => setLogin(!login)}>or register</span>
      </>
    );
  }

  function registerButton() {
    return (
      <>
        <FormButton id="register" value="Register"/>
        <span className="hover:cursor-pointer" onClick={() => setLogin(!login)}>or login</span>
      </>
    )
  }

  function forms() {
    return (
      <form
        onSubmit={(e) => e.preventDefault()}
        className="flex flex-col items-center ml-5 mr-5"
      >
        <label htmlFor="email">Email</label>
        <FormInput type="email" name="email" id="email"/>
        {!login &&
            <>
                <label htmlFor="name">Name</label>
                <FormInput type="text" name="name" id="name"/>
            </>
        }
        <label htmlFor="password">Password</label>
        <FormInput type="password" name="password" id="password"/>
        {!login &&
            <>
                <label htmlFor="passwordRepeat">Repeat Password</label>
                <FormInput type="password" name="passwordRepeat" id="passwordRepeat"/>
            </>
        }
        <div className="flex gap-2 mt-5 mb-5 items-center">
          {login ? loginButton() : registerButton()}
        </div>
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
        className="p-2 rounded-md bg-white mb-2"
        required={true}
      />
    );
  }

  function FormButton({id, value}: { id: string; value: string }) {
    return (
      <input
        type="submit"
        id={id}
        value={value}
        className="p-2 rounded-md bg-white hover:cursor-pointer transition ease-in-out delay-100 hover:-translate-y-0.5 hover:scale-105 hover:bg-gray-100 duration-200"
        onClick={() => {
          if (login) {
            // TODO: login logic
          } else {
            registerLogic().then();
          }
        }}
      />
    );
  }

  async function registerLogic() {
    let email = document.querySelector("#email") as HTMLInputElement | null;
    let name = document.querySelector("#name") as HTMLInputElement | null;
    let password = document.querySelector("#password") as HTMLInputElement | null;
    let passwordRepeat = document.querySelector("#passwordRepeat") as HTMLInputElement | null;
    if (email?.value === "" || name?.value === "" || password?.value === "" || passwordRepeat?.value === "") {
      return
    }
    if (password?.value === passwordRepeat?.value) {
      let url = process.env.NEXT_PUBLIC_ENVIRONMENT === "production" ? "https://app.kioku.dev/api/register" : "http://localhost:3001/api/register";
      const response = await fetch(url, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          email: email?.value,
          name: name?.value,
          password: password?.value
        }),
      });
      if (response.ok) {
        toast.info("Account was created!", {toastId: "accountToast"});
        setLogin(true);
      } else {
        toast.error("Account already exists!", {toastId: "accountToast"});
      }
    } else {
      toast.error("The passwords do not match!", {toastId: "passwordToast"});
    }
  }
}
