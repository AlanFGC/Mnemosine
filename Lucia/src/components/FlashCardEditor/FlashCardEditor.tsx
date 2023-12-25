import { useCallback, useState } from 'react';
import { Answer } from '../../Data/FlashcardData/Answer/Answer';
import AnswerEditor from '../AnswerEditor/AnswerEditor';
import FlashCardContentEditor from '../FlashCardContentEditor/FlashCardContentEditor';

type FlashCardEditorProps = {
  flashcardID: string | null;
};

// TODO set up this line to load a flashcard from the backend given and id.
// eslint-disable-next-line @typescript-eslint/no-unused-vars
function FlashCardEditor({ flashcardID }: FlashCardEditorProps) {
  // Raw input
  const [, setCardContent] = useState('');
  // answer tokens
  const [answerMap, setAnswerMap] = useState <Map<number, Answer>>(new Map());
  const handleEditorContent = (content: string) => {
    setCardContent(content);
  };

  const addToken = (token: number, anwser: Answer) => {
    setAnswerMap(new Map(answerMap.set(token, anwser)));
  };

  const removeToken = (key: number) => {
    const newTokens = new Map(answerMap);
    newTokens.delete(key);
    setAnswerMap(newTokens);
  };

  const editTokenValue = useCallback((key: number, newAnswer: Answer) => {
    if (answerMap.has(key)) {
      setAnswerMap(new Map(answerMap.set(key, newAnswer)));
    }
  }, [answerMap]);

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
        editTokenValue={editTokenValue}
      />
    </div>
  );
}

export default FlashCardEditor;
