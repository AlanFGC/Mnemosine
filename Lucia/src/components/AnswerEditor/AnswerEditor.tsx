import { Select, Space } from 'antd';
import { useCallback, useState } from 'react';
import { QuestionType } from '../../abstractions/flashcardsAbstractions/Answer/Answer';
import SingleAnswer from './SingleAnswer/SingleAnswer';

interface AnswerEditorProps {
  field: number;
}

export default function AnswerEditor({ field }: AnswerEditorProps) {
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
          style={{ width: 120 }}
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
        {questionType === QuestionType.Open && <h1>Open answer</h1>}
        {questionType === QuestionType.MultipleChoice && <h1>Multiple Choice Answer answer</h1>}
        {questionType === QuestionType.SingleAnswer && <SingleAnswer />}
      </div>
    </div>

  );
}
