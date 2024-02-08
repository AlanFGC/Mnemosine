import { useEffect, useState } from 'react';
import Loader from '../UtilityUIcomponents/Loader/Loader';
import FlashCardContextProvider from '../Contexts/FlashCardEditorContext';
import { EditorState } from '../Reducers/FlashcardEditorReducer';
import AnswerEditor from '../AnswerEditor/AnswerEditor';
import FlashCardContentEditor from '../FlashCardContentEditor/FlashCardContentEditor';
import FlashCardEditorCSS from './FlashCardEdtiorCSS.module.css';

type FlashCardEditorProps = {
  flashcardID: string | undefined;
};

function FlashCardEditor({ flashcardID = undefined }: FlashCardEditorProps) {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [fetchedState, setFetchedState] = useState<undefined | EditorState>(undefined);
  const [isLoading, setLoading] = useState(true);
  useEffect(
    () => {
      if (flashcardID) {
        console.log('Loading flashcard');
        // TODO FETCH FROM SERVER
      } else {
        setLoading(false);
      }
    },
    [flashcardID, isLoading, setLoading],
  );

  const contentEditorId = flashcardID ? `FLSCRDEDIT:${flashcardID}` : 'FLSCRDEDIT:NEWFLASH';

  return (
    isLoading ? <Loader />
      : (
        <FlashCardContextProvider propEditorState={fetchedState}>
          <div className={FlashCardEditorCSS.container}>
            <div className={FlashCardEditorCSS.content}>
              <FlashCardContentEditor id={contentEditorId} />
            </div>
            <div className={FlashCardEditorCSS.answer}>
              <AnswerEditor prompt="Type your answer" />
            </div>
          </div>
        </FlashCardContextProvider>
      )
  );
}

export default FlashCardEditor;
