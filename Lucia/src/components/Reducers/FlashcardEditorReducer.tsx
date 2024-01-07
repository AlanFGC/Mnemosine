import { Answer } from '../../Data/FlashcardData/Answer/Answer';

export type EditorState = {
  content: string;
  answerMap: Map<number, Answer>;
};

export const InitialState: EditorState = {
  content: '',
  answerMap: new Map<number, Answer>(),
};

type AddTokenAction = {
  type: 'addToken';
  payload: {
    key: number, answer: Answer
  };
};

type RemoveTokenAction = {
  type: 'removeToken';
  key: number;
};

type UpdateTokenAction = {
  type: 'updateToken';
  payload: {
    key: number, answer: Answer;
  }
};

type SetEditorContentAction = {
  type: 'setEditorContent';
  payload: {
    content: string
  };
};

export type FlashCardEditorActions = AddTokenAction
| RemoveTokenAction | UpdateTokenAction | SetEditorContentAction;

function FlashCardReducer(state: EditorState, action: FlashCardEditorActions): EditorState {
  const newState: EditorState = { ...state };
  switch (action.type) {
    case 'addToken':
    case 'updateToken':
      if (newState.answerMap !== null) {
        newState.answerMap.set(action.payload.key, action.payload.answer);
      }
      break;
    case 'removeToken':
      if (newState.answerMap !== null) {
        newState.answerMap.delete(action.key);
      }
      break;
    case 'setEditorContent':
      newState.content = action.payload.content;
      break;
    default: {
      throw Error('Unknown action, type not implemented');
    }
  }
  return newState;
}

export default FlashCardReducer;
