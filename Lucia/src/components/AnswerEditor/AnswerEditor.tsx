import { Select, Space } from 'antd';
import {
  useCallback, useEffect, useRef, useState,
} from 'react';
import { Answer } from '../../Data/FlashcardData/Answer/Answer';
import AnswerWrapper from './AnswerEditorAnswer/AnswerWrapper';

interface AnswerEditorProps {
  prompt: string;
  answerTokens: Map<number, Answer>;
  setAnswers: (token: number, newAnswer: Answer) => void;
}

export default function AnswerEditor({
  prompt, answerTokens, setAnswers,
}: AnswerEditorProps) {
  const [answerList, setAnswerList] = useState([...answerTokens.values()]);

  useEffect(() => {
    setAnswerList([...answerTokens.values()]);
  }, [answerTokens]);

  return (
    <div>
      <div className="topbar">
        <span>{prompt}</span>
        <Space wrap />
      </div>
      <div>
        <ul>
          {answerList.map((answer: Answer) => (
            <li>
              <AnswerWrapper answerInit={answer} setGlobalAnswer={setAnswers} />
            </li>
          ))}
        </ul>
      </div>
    </div>

  );
}
