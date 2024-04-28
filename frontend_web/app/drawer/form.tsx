"use client"

import { useForm } from 'react-hook-form';
import { z } from 'zod';

import { Button } from '@/components/ui/button';
import {
  Form,
  FormControl,
  FormDescription,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@/components/ui/form';
import { Input } from '@/components/ui/input';
import { Textarea } from '@/components/ui/textarea';
// import { toast } from '@/components/ui/use-toast';
import { zodResolver } from '@hookform/resolvers/zod';

const ACCEPTED_IMAGE_TYPES = ["image/jpeg", "image/jpg", "image/png", "image/webp"];
const FormSchema = z.object({
  bio: z
    .string()
    .min(10, {
      message: "Bio must be at least 10 characters.",
    })
    .max(160, {
      message: "Bio must not be longer than 30 characters.",
    }),
  images: z.any()
        .refine(files => {return Array.from(files).every(file => file instanceof File)}, { message: "Expected a file" })
        .refine(files => Array.from(files).every(file => ACCEPTED_IMAGE_TYPES.includes(file.type)), "Only these types are allowed .jpg, .jpeg, .png and .webp")

})

export function InputForm() {
  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      bio: "",
      images: '',
    },
  })

  function onSubmit(data: z.infer<typeof FormSchema>) {
    console.log(data);
    // toast({
    //   title: "You submitted the following values:",
    //   description: (
    //     <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
    //       <code className="text-white">{JSON.stringify(data, null, 2)}</code>
    //     </pre>
    //   ),
    // })
  }

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className=" space-y-12  flex flex-col justify-stretch px-6">
      <div className='bg-white p-10 rounded-2xl' > <FormField
                 
          control={form.control}
          name="bio"
          render={({ field }) => (
            <FormItem className = 'my-5' >  
              <FormLabel>Add Note as Text</FormLabel>
              <FormControl>
                <Textarea
                  placeholder="Tell us a little bit about yourself"
                  className="h-50 rounded "
                  {...field}

                />
              </FormControl>
              <FormDescription>
                  

              </FormDescription>
              <FormMessage  />
            </FormItem>
          )}
        />

        <FormField
                    control={form.control}
                    name="images"
                    render={({ field: { value, onChange, ...fieldProps } }) => (
                      <FormItem>
                        <FormLabel>Add Note As a Picture</FormLabel>
                        <FormControl>
                          <Input
                            {...fieldProps}
                            placeholder="Picture"
                            type="file"
                            accept="image/*, application/pdf"
                            className='w-2/3'
                            onChange={(event) =>
                              onChange(event.target.files && event.target.files[0])
                              
                            }/>
                            </FormControl>
                            <FormDescription>
                              
                            </FormDescription>
                            <FormMessage />
                        </FormItem>
                    )}
                />
      <div className='flex justify-end mt-10'> <Button className = 'bg-slate-300 hover:bg-slate-200 rounded' type="submit" >Submit</Button></div>
        </div>
      </form>
    </Form>
  )
}
