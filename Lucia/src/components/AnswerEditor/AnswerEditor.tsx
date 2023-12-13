import { Space } from 'antd';
import { Answer } from '../../Data/FlashcardData/Answer/Answer';
import AnswerWrapper from './AnswerEditorAnswer/AnswerWrapper';

interface AnswerEditorProps {
  prompt: string;
  answerTokens: Map<number, Answer>;
  editTokenValue: (token: number, newAnswer: Answer) => void;
}

export default function AnswerEditor({
  prompt, answerTokens, editTokenValue,
}: AnswerEditorProps) {
  const validKeys = ['field', 'answers', 'incorrectAnswers', 'explanation', 'questionType'];
  const isValidKey = (key: string) => validKeys.includes(key);

  const updateAnswer = (field: number, propertyName: string, newValue: string) => {
    const oldAnswer = answerTokens.get(field);
    if (oldAnswer && isValidKey(propertyName)) {
      const updatedAnswer = { ...oldAnswer, [propertyName]: newValue };
      editTokenValue(field, updatedAnswer);
    }
  };
  return (
    <div>
      <div className="topbar">
        <span>{prompt}</span>
        <Space wrap />
      </div>
      <div>
        <ul>
          {[...answerTokens.values()].map((answer: Answer) => (
            <li>
              <AnswerWrapper answerInit={answer} updateAnswer={updateAnswer} />
            </li>
          ))}
        </ul>
      </div>
    </div>

  );
}
