import { QuestionType, Answer } from './Answer';

export default class AnswerBuilder {
  private field: number = 0;

  private answers: string[] = [];

  private incorrectAnswers: string[] = [];

  private explanation: string = '';

  private questionType: QuestionType = QuestionType.Undefined;

  setField(field: number): AnswerBuilder {
    this.field = field;
    return this;
  }

  setAnswers(answers: string[]): AnswerBuilder {
    this.answers = answers;
    return this;
  }

  setIncorrectAnswers(incorrectAnswers: string[]): AnswerBuilder {
    this.incorrectAnswers = incorrectAnswers;
    return this;
  }

  setExplanation(explanation: string): AnswerBuilder {
    this.explanation = explanation;
    return this;
  }

  setQuestionType(questionType: QuestionType): AnswerBuilder {
    this.questionType = questionType;
    return this;
  }

  build(): Answer {
    return {
      field: this.field,
      answers: this.answers,
      incorrectAnswers: this.incorrectAnswers,
      explanation: this.explanation,
      questionType: this.questionType,
    };
  }
}
