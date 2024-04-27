"use client"
import { useSignUpMutation } from '@/lib/features/auth/authApi'; // Check and verify this import path
import React, { useState } from 'react';
import { Button } from '../ui/button';
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from '../ui/form';
import { Input } from '../ui/input';
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import SuccessMessage from '../messages/successMessage';
import ErrorMessage from '../messages/errorMessage';



const formSchema = z.object({
  username: z.string().min(2, {
    message: "Username must be at least 2 characters.",
  }),
  email:z.string().min(2, {message:"email must be at least 2 characters"}),
  role:z.string().min(2, {message:"Role must be at least 2 characters"}),
  password:z.string().min(2, {message:"Password must be at least 2 characters"}),
 
})
const RegisterComponent = () => {

    const [signUp, { isError, isLoading, isSuccess, data, error }] = useSignUpMutation();


    const form = useForm<z.infer<typeof formSchema>>({
      resolver: zodResolver(formSchema),
      defaultValues: {
        username: "",
        email:"",
        password:"",
        role:"",
      },
    })

    const onSubmit = async (values: z.infer<typeof formSchema>) => {

        console.log("iserror: " + isError, error, data)

     
        await signUp({
          username: values.username,
          email:values.email,
          password:values.password,
          role:values.role,
        });
        console.log("iserror: " + isError, error, data)

       
    };
    
    if(data&&data?.accessToken){
      localStorage.setItem("accessToken", data.accessToken)
    }
   
    return (
        <div className='flex min-h-screen w-full justify-center items-center'>
          <div className="w-max p-4 shadow-md bg-white rounded-lg">
          <Form {...form} >
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
       

<FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Username</FormLabel>
              <FormControl>
                <Input placeholder="username" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />





<FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Email</FormLabel>
              <FormControl>
                <Input placeholder="email" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />



<FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Password</FormLabel>
              <FormControl>
                <Input placeholder="password" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />


<FormField
          control={form.control}
          name="role"
          render={({ field }) => (
            <FormItem>
              <FormLabel>Role</FormLabel>
              <FormControl>
                <Input placeholder="role" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />


{isSuccess?<SuccessMessage message='Successfully Registered'/>:""}
{isError?<ErrorMessage message='Error Ocurred in  Registered'/>:""}
   <div className="flex justify-end">
       
   <Button type="submit" disabled={isLoading}>{isLoading?"Loading...":"Submit"}</Button>
   </div>
      </form>
    </Form>
          </div>
        </div>
    );
};

export default RegisterComponent;
