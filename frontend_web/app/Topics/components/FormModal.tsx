import { useState } from 'react';

import { useRouter } from 'next/navigation';
import Modal from 'react-modal';
import {
  useDispatch,
  useSelector,
} from 'react-redux';
import { z } from 'zod';

import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
import { Textarea } from '@/components/ui/textarea';
import { createSubjectAction } from '@/lib/features/topics/topicSlice';
import { FormDataModel } from '@/types/formData';

const schema = z.object({
  name: z.string({ required_error: "Name is required" }).min(3, { message: "Name must be at least 3 characters long" }).max(20, { message: "Name must be at most 255 characters long" }).refine(value => value.trim() !== "", "Name cannot be empty"),
  description: z.string({ required_error: "Description is required" }).min(5, { message: "Description must be at least 3 characters long" }).max(50, { message: "Description must be at most 255 characters long" }).refine(value => value.trim() !== "", "Description cannot be empty"),
})

interface Props {
  modalIsOpen: boolean;
  closeModal: () => void;
}


export const FormModal = (props: Props) => {
  const [formValues, setFormValues] = useState({
    topic: "",
    description: ""
  });

  const dispatch = useDispatch();
  const {isSuccess, isLoading, errors} = useSelector((state: any) => state.subjects.topic);
  const router = useRouter();

  function onValueChange(event: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) {
    const { name, value } = event.target;
    setFormValues((prevValues) => {
      return {
        ...prevValues,
        [name]: value,
      };
    });
  }

  const handleCloseForm = () => {
    setFormValues({
      topic: "",
      description: ""
    });

    props.closeModal()
  }

  async function onSubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const formData: FormDataModel = { subjectName: formValues.topic, description: formValues.description }

    dispatch(createSubjectAction(formData));
  }

  if (isSuccess) {
    router.push('/Topics/1')
  }

  return <Modal isOpen={props.modalIsOpen}
    onRequestClose={() => { props.closeModal() }}
    className="fixed inset-0 flex items-center justify-center bg-transparent"
    shouldCloseOnOverlayClick={true}>
    <Card className="w-[400px] max-w[400px] h-[400px] items-center">
      <form onSubmit={onSubmit}>
        <CardHeader className='py-8'>
          <CardTitle>Create Topic</CardTitle>
        </CardHeader>
        <CardContent>
          <div className="grid w-full items-center gap-10">
            <div className="flex flex-col space-y-3">
              <Label htmlFor="topic">Topic name</Label>
              <Input id="topic" name="topic" placeholder="Name of your project" onChange={onValueChange} />
            </div>
            <div className="flex flex-col space-y-3">
              <Label htmlFor="description">Topic description</Label>
              <Textarea id="description" name="description" placeholder="Name of your project" onChange={onValueChange} />
            </div>
          </div>
        </CardContent>
        <CardFooter className="flex justify-around">
          <Button variant="outline" onClick={handleCloseForm}>Cancel</Button>
          <Button className=' bg-mainColor' type='submit' disabled={isLoading}>Create</Button>
        </CardFooter>
      </form>
    </Card>
  </Modal>
}