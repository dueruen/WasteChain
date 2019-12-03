import React, { Fragment } from 'react';
// import './App.css';

import { ApolloProvider, useQuery } from '@apollo/react-hooks';
import gqlclient from './api/gql/ApolloClient'
import Pages from './pages';


const App = () => {
    return (
        <ApolloProvider client={gqlclient}>
            {/* <IsLoggedIn /> */}
            <Pages></Pages>
        </ApolloProvider>
    )
}

export default App;
