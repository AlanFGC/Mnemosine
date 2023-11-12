import { Select, Space } from 'antd';
import { useState } from 'react';
import { QuestionType } from '../../abstractions/flashcardsAbstractions/Answer/Answer';

interface AnswerEditorProps {
  field: number;
}

function handleChange(): void {

}

export default function AnswerEditor({ field }: AnswerEditorProps) {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [myValue, setMyValue] = useState(QuestionType.Open);
  const questionValues = Object.values(QuestionType);

  return (
    <div>
      <div className="topbar">
        <span>{field}</span>
        <Space wrap />
        <Select
          defaultValue="Question type"
          style={{ width: 120 }}
          onChange={handleChange}
          options={[
            {
              ...questionValues.map((questionType: QuestionType) => (
                { value: questionType, label: questionType, disabled: false }
              )),
            },
          ]}
        />
      </div>
    </div>

  );
}
