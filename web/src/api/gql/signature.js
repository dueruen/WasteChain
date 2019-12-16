import gql from "graphql-tag";
import { useMutation } from '@apollo/react-hooks';

/**
 * Mutation to continue double sign process
 */
const CONTINUE_DOUBLE_SIGN =
    gql`
    mutation ContinueDoubleSign($continueID: String!, $newHolderID: String!, $newHolderPassword: String!)
    {continueDoubleSign(request: {continueID: $continueID, newHolderID: $newHolderID, newHolderPassword: $newHolderPassword})}
`;


