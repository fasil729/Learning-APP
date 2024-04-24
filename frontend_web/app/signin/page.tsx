import Head from 'next/head';
import React from 'react';
import Link from 'next/link';
import { FaFacebookF, FaLinkedinIn, FaGoogle, FaEnvelope } from 'react-icons/fa';

export default function SignIn() {
    return (
        <div className="flex flex-col items-center justify-center min-h-screen py-2 bg-gray-200">
            <Head>
                <title>Sign In</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <main className="flex flex-col items-center justify-center w-full flex-1 px-4 md:px-20 text-center"> {/* Added responsive padding */}
                <div className='bg-white rounded-2xl shadow-2xl flex flex-col md:flex-row w-full md:w-2/3 max-w-4xl'>
                    <div className='w-full md:w-3/5 p-5'>
                        <div className='text-left font-bold'>
                            <span className="text-blue-500">Learning App</span>
                        </div>
                        <div className='py-10'>
                            <h2 className='text-3xl font-bold text-blue-500 mb-2'>Sign in to Learning App</h2>
                            <div className='border-2 w-10 border-blue-500 inline-block mb-2'></div>
                            <div className='flex justify-center my-2'>
                                <a href="#" className='border-2 border-gray-200 rounded-full p-3 mx-1'>
                                    <FaFacebookF className='text-blue-500' />
                                </a>
                                <a href="#" className='border-2 border-gray-200 rounded-full p-3 mx-1'>
                                    <FaLinkedinIn className='text-blue-500' />
                                </a>
                                <a href="#" className='border-2 border-gray-200 rounded-full p-3 mx-1'>
                                    <FaGoogle className='text-blue-500' />
                                </a>
                            </div>
                            <p className='text-gray-500 my-3'>Or use your email account</p>
                            <div className='flex flex-col items-center'>
                                <div className='bg-gray-100 w-full md:w-64 p-2 flex items-center rounded-full mb-3'>
                                    <FaEnvelope className='text-gray-400 mr-2' />
                                    <input type='email' name='email' placeholder='Email' className='bg-gray-100 outline-none text-sm flex-1' />
                                </div>
                                <div className='bg-gray-100 w-full md:w-64 p-2 flex items-center rounded-full mb-3'>
                                    <FaEnvelope className='text-gray-400 mr-2' />
                                    <input type='password' name='password' placeholder='password' className='bg-gray-100 outline-none text-sm flex-1' />
                                </div>
                                <div  className='flex w-full md:w-64 mb-5 ' >
                                    <label htmlFor="" className='flex items-center text-xs'><input type='checkbox' name='remember'/> Remember me </label>
                                </div>
                                <div className="flex justify-center">
                                    <button className="bg-blue-500 text-white px-6 py-2 rounded-md font-semibold hover:bg-blue-600">Sign In</button>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="w-full md:w-2/5 bg-blue-500 text-white rounded-tr-2xl rounded-br-2xl py-8 md:py-36 px-6 md:px-12">
                        <div className="text-2xl md:text-3xl font-bold mb-2">Welcome to Learning App!</div>
                        <div className="">
                            <p className="mb-4 md:mb-6">Fill up your personal information to start your learning journey with us.</p>
                            <p className="mb-2 md:mb-4">Already have an account?</p>
                            <Link href="/register">
                                <button className='border-2 border-white rounded-full px-8 py-2 inline-block font-semibold hover:bg-white hover:text-blue-500'>Sign Up</button>
                            </Link>
                        </div>
                    </div>
                </div>
            </main>
        </div>
    );
}
