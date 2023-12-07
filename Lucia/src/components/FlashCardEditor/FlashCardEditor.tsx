import { useState } from 'react';
import { Answer, AnswerToken } from '../../Data/FlashcardData/Answer/Answer';
import AnswerEditor from '../AnswerEditor/AnswerEditor';
import FlashCardContentEditor from '../FlashCardContentEditor/FlashCardContentEditor';

function FlashCardEditor() {
  // Raw input
  const [, setCardContent] = useState('');
  // answer tokens
  const [answerMap, setAnswerMap] = useState <Map<number, Answer>>(new Map());
  const handleEditorContent = (content: string) => {
    setCardContent(content);
  };

  /*
  Set token for update on start
  const setToken = (token: number, answer: Answer) => {
    setAnswerMap(new Map(answerMap.set(token, answer)));
  };
  */

  const addToken = (token: number, anwser: Answer) => {
    setAnswerMap(new Map(answerMap.set(token, anwser)));
  };

  const removeToken = (key: number) => {
    const newTokens = new Map(answerMap);
    newTokens.delete(key);
    setAnswerMap(newTokens);
  };

  const editTokenKey = (key: number, newAnswer: Answer) => {
    if (answerMap.has(key)) {
      setAnswerMap(new Map(answerMap.set(key, newAnswer)));
    }
  };

  return (
    <div>
      <FlashCardContentEditor
        setCardContent={handleEditorContent}
        addAnswerToken={addToken}
        removeAnswerToken={removeToken}
      />
      <AnswerEditor
        prompt="Add your answers here :)"
        answerTokens={answerMap}
        setAnswers={editTokenKey}
      />
    </div>
  );
}

export default FlashCardEditor;
