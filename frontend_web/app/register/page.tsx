import Head from 'next/head';
import Link from 'next/link';
import React from 'react';
import { FaFacebookF, FaLinkedinIn, FaGoogle, FaUser, FaUserAlt, FaEnvelope, FaLock } from 'react-icons/fa';

export default function SignIn() {

    function Register() {
        const { push } = useRouter();
        const [register, { isLoading, error, success, isError }] = useCreateAccountMutation();
        const dispatch = useDispatch();
        const nameRef = useRef<HTMLInputElement>(null);
        const usernameRef = useRef<HTMLInputElement>(null);
        const emailRef = useRef<HTMLInputElement>(null);
        const passwordRef = useRef<HTMLInputElement>(null);
      
        const submitForm = async (e: React.FormEvent<HTMLFormElement>) => {
          e.preventDefault();
      
          const name = nameRef.current?.value;
          const username = usernameRef.current?.value;
          const email = emailRef.current?.value;
          const password = passwordRef.current?.value;
      
          try {
            const response = await register({ name, username, email, password }).unwrap();
      
            if (response?.accessToken) {
              dispatch(setCredentials({ response }));
              toast.success("Successfully registered and logged in.");
              push('/dashboard/home');
            }
      
          } catch (err) {
            console.error('Registration failed:', err);
          }
        };
      




    return (
        <div className="flex flex-col items-center justify-center min-h-screen py-2 bg-gray-200">
            <Head>
                <title>Sign In</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <main className="flex flex-col items-center justify-center w-full flex-1 px-20 text-center">
                <div className="bg-white rounded-2xl shadow-2xl flex w-2/3 max-w-4xl">
                    {/* Left Panel */}
                    <div className="w-3/5 p-5">
                        <div className="text-left font-bold">
                            <span className="text-blue-500">Learning App</span>
                        </div>
                        <div className="py-10">
                            <h2 className="text-3xl font-bold text-blue-500 mb-2">Sign in to Learning App</h2>
                            <div className="border-2 w-10 border-blue-500 inline-block mb-2"></div>
                            <div className="flex justify-center my-2">
                                <a href="#" className="border-2 border-gray-200 rounded-full p-3 mx-1">
                                    <FaFacebookF className="text-blue-500" />
                                </a>
                                <a href="#" className="border-2 border-gray-200 rounded-full p-3 mx-1">
                                    <FaLinkedinIn className="text-blue-500" />
                                </a>
                                <a href="#" className="border-2 border-gray-200 rounded-full p-3 mx-1">
                                    <FaGoogle className="text-blue-500" />
                                </a>
                            </div>
                            <p className="text-gray-500 my-3">Or use your email account</p>
                            {/* Sign In Form */}
                            <div className="py-10">
                                <FormField label="" type="text" id="name" placeholder="Enter your name" icon={<FaUser />} />
                                <FormField label="" type="text" id="username" placeholder="Enter your username" icon={<FaUserAlt />} />
                                <FormField label="" type="email" id="email" placeholder="Enter your email" icon={<FaEnvelope />} />
                                <FormField label="" type="password" id="password" placeholder="Enter your password" icon={<FaLock />} />
                                <div className="flex justify-center">
                                    <button className="bg-blue-500 text-white px-6 py-2 rounded-md font-semibold hover:bg-green-600">Sign Up</button>
                                </div>
                            </div>
                        </div>
                    </div>
                    {/* Right Panel */}
                    <div className="w-2/5 bg-blue-500 text-white rounded-tr-2xl rounded-br-2xl py-36 px-12">
                        <div className="text-3xl font-bold mb-2">Welcome to the Learning App!</div>
                        <div className="">
                            <p className="mb-6">Fill up your personal information to start your learning journey with us.</p>
                            <p className="mb-4">Already have an account?</p>
                            <Link href="/signin">
                            <button className='border-2 border-white rounded-full px-12 py-2 inline-block font-semibold hover:bg-white hover:text-blue-500'>Sign In</button>
                        </Link>                        </div>
                    </div>
                </div>
            </main>
        </div>
    );
}

// Reusable form field component
function FormField({ label, type, id, placeholder, icon }: { label: string, type: string, id: string, placeholder: string, icon: React.ReactNode }) {
    return (
        <div className="flex flex-col mb-3">
            <label htmlFor={id} className="text-gray-600 mb-1">{label}</label>
            <div className="bg-gray-100 flex items-center rounded-md">
                <div className="text-gray-400 ml-3">{icon}</div>
                <input type={type} id={id} name={id} placeholder={placeholder} className="bg-gray-100 outline-none px-3 py-2 flex-1 text-sm" />
            </div>
        </div>
    );
}
