import { Select, Space } from 'antd';
import { useState } from 'react';
import { Answer, QuestionType } from '../../../Data/FlashcardData/Answer/Answer';
import SingleAnswer from './SingleAnswer/SingleAnswer';
import MultipleChoice from './MultipleChoice/MultipleChoice';
import OpenAnswer from './OpenAnswer/OpenAnswer';

interface AnswerEditorProps {
  answerInit: Answer;
  updateAnswer: (field: number, propertyName: string, newValue: string) => void;
}

export default function AnswerWrapper({ answerInit, updateAnswer }: AnswerEditorProps) {
  const [questionType, setQuestionType] = useState(answerInit.questionType);

  const handleQuestionChange = (value: QuestionType) => {
    setQuestionType(value);
    updateAnswer(answerInit.field, 'questionType', value);
  };

  const handleInputChange = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
    const { name, value } = event.target;
    updateAnswer(answerInit.field, name, value);
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
      </div>
      <div>
        {questionType === QuestionType.Open
         && <OpenAnswer handleInputChange={handleInputChange} />}
        {/* {questionType === QuestionType.MultipleChoice
        && <MultipleChoice handleInputChange={handleInputChange} />} */}
        {questionType === QuestionType.SingleAnswer
         && <SingleAnswer handleInputChange={handleInputChange} />}
      </div>
    </div>

  );
}
