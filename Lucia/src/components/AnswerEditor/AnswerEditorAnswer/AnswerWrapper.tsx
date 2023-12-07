import { Select, Space } from 'antd';
import { useCallback, useState } from 'react';
import { Answer, QuestionType } from '../../../Data/FlashcardData/Answer/Answer';
import SingleAnswer from './SingleAnswer/SingleAnswer';
import MultipleChoice from './MultipleChoice/MultipleChoice';
import OpenAnswer from './OpenAnswer/OpenAnswer';

interface AnswerEditorProps {
  answerInit: Answer;
  setGlobalAnswer: (token: number, newAnswer: Answer) => void;
}

export default function AnswerWrapper({ answerInit, setGlobalAnswer }: AnswerEditorProps) {
  const [questionType, setQuestionType] = useState(QuestionType.Open);

  const [answer, setAnswer] = useState({ ...answerInit });

  const handleQuestionChange = useCallback((value: QuestionType) => {
    setQuestionType(value);
    setGlobalAnswer(answer.field, { ...answer, questionType: value });
  }, [setQuestionType, answer]);


  const handleAnswerChange = useCallback((newAnswer: Answer) => {
    setAnswer(newAnswer);
    setGlobalAnswer(answer.field, answer);
  }, [setAnswer, setGlobalAnswer]);

  return (
    <div>
      <div className="topbar">
        <span>{answer.field}</span>
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
        {questionType === QuestionType.Open && <OpenAnswer onChange={ handleAnswerChange } />}
        {questionType === QuestionType.MultipleChoice && <MultipleChoice onChange={ handleAnswerChange } />}
        {questionType === QuestionType.SingleAnswer && <SingleAnswer onChange={ handleAnswerChange } />}
      </div>
    </div>

  );
}
