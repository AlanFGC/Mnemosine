import { Button } from 'antd';
import { useContext } from 'react';
import AnswerWrapper from './AnswerEditorAnswer/AnswerWrapper';
import { FlashCardEditorContext } from '../Contexts/FlashCardEditorContext';
import { DefaultAnswer } from '../../Data/FlashcardData/Answer/Answer';

type AnswerEditorProps = {
  prompt: string;
};

/* const VALIDKEYS = ['field', 'answers', 'incorrectAnswers', 'explanation', 'questionType']; */

export default function AnswerEditor({ prompt }: AnswerEditorProps) {
  const { state: editorState, dispatch } = useContext(FlashCardEditorContext);

  const addToken = () => {
    const newKey = Math.max(...editorState.answerMap.keys(), 0) + 1;
    dispatch({
      type: 'addToken',
      payload: {
        key: newKey,
        answer: { ...DefaultAnswer, field: newKey },
      },
    });
  };

  return (
    <div>
      <div className="topbar">
        <span>{prompt}</span>
      </div>
      <div>
        <ul>
          {
        Array.from(editorState.answerMap.values()).map((answer) => (
          <li key={answer.field}>
            <AnswerWrapper answerInit={answer} />
          </li>
        ))
      }
        </ul>
        <Button onClick={addToken}>Add Answer</Button>
      </div>
    </div>

  );
}
