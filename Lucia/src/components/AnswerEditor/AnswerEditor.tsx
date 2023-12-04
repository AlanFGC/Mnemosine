import { Select, Space } from 'antd';
import { useCallback, useState } from 'react';
import { Answer, QuestionType } from '../../Data/FlashcardData/Answer/Answer';
import MultipleChoice from './AnswerEditorAnswer/MultipleChoice/MultipleChoice';
import OpenAnswer from './AnswerEditorAnswer/OpenAnswer/OpenAnswer';
import SingleAnswer from './AnswerEditorAnswer/SingleAnswer/SingleAnswer';

interface AnswerEditorProps {
  prompt: string;
  answerTokens: Map<number, Answer>;
  setAnswers: (token: number, newAnswer: Answer) => void;
}

export default function AnswerEditor({ prompt, answerTokens, setAnswers }: AnswerEditorProps) {
  const [questionType, setQuestionType] = useState(QuestionType.Open);

  const handleChange = useCallback((value: QuestionType) => {
    setQuestionType(value);
  }, [setQuestionType]);

  return (
    <div>
      <div className="topbar">
        <span>{prompt}</span>
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
