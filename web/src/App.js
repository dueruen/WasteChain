import React from 'react';

import { ApolloProvider } from '@apollo/react-hooks';
import gqlclient from './api/gql/ApolloClient'
import Pages from './pages';


const App = () => {
    return (
        <ApolloProvider client={gqlclient}>
            <Pages></Pages>
        </ApolloProvider>
    )
}

export default App;
