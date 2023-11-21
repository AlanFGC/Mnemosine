export interface Answer {
  field: number;
  answers: string[];
  incorrectAnswers: string[];
  explanation: string;
  questionType: QuestionType;
}

export enum QuestionType {
  Open = 'open',
  SingleAnswer = 'single',
  MultipleChoice = 'multiple',
  Undefined = 'undefined',
}
