import { Button, Select, Space } from 'antd';
import { useState } from 'react';
import { Answer, QuestionType } from '../../../Data/FlashcardData/Answer/Answer';
import SingleAnswer from './SingleAnswer/SingleAnswer';
import MultipleChoice from './MultipleChoice/MultipleChoice';
import OpenAnswer from './OpenAnswer/OpenAnswer';
import { ANSWERS, EXPLANATION, INCORRECTANSWERS } from './dtypes';

interface AnswerEditorProps {
  answerInit: Answer;
}

export default function AnswerWrapper({ answerInit }: AnswerEditorProps) {
  const [answer, setAnswer] = useState<Answer>(answerInit);

  const handleQuestionChange = (value: QuestionType) => {
    setAnswer({ ...answer, questionType: value });
  };

  const handleInputChange = (name: string, value:string[]) => {
    switch (name) {
      case EXPLANATION:
        setAnswer({ ...answer, explanation: value[0] });
        break;
      case ANSWERS:
        setAnswer({ ...answer, answers: value });
        break;
      case INCORRECTANSWERS:
        setAnswer({ ...answer, incorrectAnswers: value });
        break;
      default:
        throw new Error(`Failed to parse input change on answer ${name}`);
    }
  };

  const onSave = () => {
  };

  return (
    <div>
      <div className="topbar">
        <Space wrap />
        <Select
          defaultValue={null}
          style={{}}
          onChange={handleQuestionChange}
          options={[

            {
              value: QuestionType.MultipleChoice,
              label: QuestionType.MultipleChoice,
              disabled: false,
            },
            {
              value: QuestionType.SingleAnswer,
              label: QuestionType.SingleAnswer,
              disabled: false,
            },
            {
              value: QuestionType.Open,
              label: QuestionType.Open,
              disabled: false,
            },

          ]}
        />
        <Button onClick={onSave}>Save</Button>
      </div>
      <div>
        {answer.questionType === QuestionType.Open
         && <OpenAnswer handleInputChange={handleInputChange} />}
        {answer.questionType === QuestionType.MultipleChoice
        && <MultipleChoice handleInputChange={handleInputChange} />}
        {answer.questionType === QuestionType.SingleAnswer
         && <SingleAnswer handleInputChange={handleInputChange} />}
      </div>
    </div>

  );
}
