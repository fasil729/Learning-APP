"use client";
import React from 'react';

import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { useForm } from 'react-hook-form';
import {
  useDispatch,
  useSelector,
} from 'react-redux';
import { z } from 'zod';

import { signInAction } from '@/lib/features/auth/authSlice';
import { User } from '@/types/user';
import { zodResolver } from '@hookform/resolvers/zod';

import ErrorMessage from '../messages/errorMessage';
import SuccessMessage from '../messages/successMessage';
import {
  Form,
  FormField,
} from '../ui/form';

const formSchema = z.object({
  username: z.string().min(2, {
    message: "Username must be at least 2 characters.",
  }),
  password: z
    .string()
    .min(2, { message: "Password must be at least 2 characters" }),
});

const LoginComponent = () => {
  const dispatch = useDispatch();
  const router = useRouter();
  var { isLoading, isAuthenticated, errors } = useSelector((state: any) => state.auth.data);

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    var user: User = {username: values.username, password: values.password}
    dispatch(signInAction(user));
  };

  if (isAuthenticated) {
    router.push("/Topics")
  }

  return (
    <div className="bg-white dark:bg-gray-900">
      <div className="flex justify-center h-screen">
        <div
          className="hidden b bg-cover lg:block lg:w-2/3"
          style={{ backgroundImage: "url(/images2.jpeg)" }}
        >
          <div className="flex items-center h-full px-20 bg-gray-900 bg-opacity-40">
            <div>
              {/* <h2 className="text-[40px] font-bold text-white sm:text-3xl">
                 Login In To Brilliant Wep App
              </h2> */}
              <p className="max-w-xl mt-3 text-[20px]  text-gray-300">
              Login In To Brilliant Wep App and Explore The World Of Learning. Start your journey today and
                unlock endless possibilities for growth and development.
              </p>
            </div>
          </div>
        </div>
        <div className="flex items-center w-full max-w-md px-6 mx-auto lg:w-2/6">
          <div className="flex-1">
            <div className="text-center">
              <div className="flex justify-center mx-auto">
                <img
                  className="w-auto h-7 sm:h-8"
                  src="https://merakiui.com/images/logo.svg"
                  alt=""
                />
              </div>
             
            </div>
            <div className="mt-8">
              <Form {...form}>
                <form
                  onSubmit={form.handleSubmit(onSubmit)}
                  className="space-y-8"
                >
                  <FormField
                    control={form.control}
                    name="username"
                    render={({ field }) => (
                      <div>
                        <label
                          htmlFor="username"
                          className="block mb-2 text-sm text-gray-600 dark:text-gray-200"
                        >
                          Username
                        </label>
                        <input
                          type="text"
                          id="username"
                          placeholder="Username"
                          className="block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-lg dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40"
                          {...field}
                        />
                      </div>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="password"
                    render={({ field }) => (
                      <div>
                        <label
                          htmlFor="password"
                          className="block mb-2 text-sm text-gray-600 dark:text-gray-200"
                        >
                          Password
                        </label>
                        <input
                          type="password"
                          id="password"
                          placeholder="Your Password"
                          className="block w-full px-4 py-2 mt-2 text-gray-700 placeholder-gray-400 bg-white border border-gray-200 rounded-lg dark:placeholder-gray-600 dark:bg-gray-900 dark:text-gray-300 dark:border-gray-700 focus:border-blue-400 dark:focus:border-blue-400 focus:ring-blue-400 focus:outline-none focus:ring focus:ring-opacity-40"
                          {...field}
                        />
                      </div>
                    )}
                  />
                  {isAuthenticated ? (
                    <SuccessMessage message="Logged In Successfully" />
                  ) : (
                    ""
                  )}
                  {errors ? (
                    <ErrorMessage message="Incorrect username or password" />
                  ) : (
                    ""
                  )}
                  <div className="mt-6">
                    <button
                      type="submit"
                      className="w-full px-4 py-2 tracking-wide text-white transition-colors duration-300 transform bg-blue-500 rounded-lg hover:bg-blue-400 focus:outline-none focus:bg-blue-400 focus:ring focus:ring-blue-300 focus:ring-opacity-50"
                      disabled={isLoading}
                    >
                      {isLoading ? "Loading..." : "Sign in"}
                    </button>
                  </div>
                </form>
              </Form>
              <p className="mt-6 text-sm text-center text-gray-400">
                Dont have an account yet?{" "}
                <Link href="/register">
                  <p className="text-blue-500 focus:outline-none focus:underline hover:underline">
                    Sign up
                  </p>
                </Link>
                .
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginComponent;
