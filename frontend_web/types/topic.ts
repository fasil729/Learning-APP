import { FormDataModel } from "./formData";

export interface Topic {
    subject: Subject;
    chapters: Chapter[];
}

export interface Subject {
    id: number,
    name: string,
}

export interface Chapter {
    name: string;
    lessons: string[];
}

export interface TopicState {
    topics: {
        isLoading: boolean;
        data: FormDataModel | null;
        errors: string | null;
    },
    topic: {
        isLoading: boolean;
        data: Topic | null;
        errors: string | null;
    }
}

