export interface Answer {
  field: number;
  answers: string[];
  incorrectAnswers: string[];
  explanation: string;
  questionType: QuestionType;
}

export interface AnswerToken {
  field: number;
}

export enum QuestionType {
  Open = 'open',
  SingleAnswer = 'single',
  MultipleChoice = 'multiple',
  Undefined = 'undefined',
}

export const DefaultAnswer: Answer = {
  field: 0,
  answers: [],
  incorrectAnswers: [],
  explanation: '',
  questionType: QuestionType.Open,
};
