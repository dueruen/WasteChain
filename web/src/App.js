import React, { Fragment } from 'react';
// import './App.css';

import { ApolloClient } from 'apollo-client';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { ApolloProvider, useQuery } from '@apollo/react-hooks';
import { HttpLink } from 'apollo-link-http';
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
