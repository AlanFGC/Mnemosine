import { Select, Space } from 'antd';
import { useCallback, useState } from 'react';
import { QuestionType } from '../../../Data/FlashcardData/Answer/Answer';
import SingleAnswer from './SingleAnswer/SingleAnswer';
import MultipleChoice from './MultipleChoice/MultipleChoice';
import OpenAnswer from './OpenAnswer/OpenAnswer';

interface AnswerEditorProps {
  field: number;
}

export default function AnswerWrapper({ field }: AnswerEditorProps) {
  const [questionType, setQuestionType] = useState(QuestionType.Open);

  const handleChange = useCallback((value: QuestionType) => {
    setQuestionType(value);
  }, [setQuestionType]);

  return (
    <div>
      <div className="topbar">
        <span>{field}</span>
        <Space wrap />
        <Select
          defaultValue={null}
          style={{}}
          onChange={handleChange}
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
        {questionType === QuestionType.Open && <OpenAnswer />}
        {questionType === QuestionType.MultipleChoice && <MultipleChoice />}
        {questionType === QuestionType.SingleAnswer && <SingleAnswer />}
      </div>
    </div>

  );
}
