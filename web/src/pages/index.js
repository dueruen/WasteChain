import React, { Fragment } from 'react';
import { Router } from '@reach/router';
import { PageContainer } from '../components';

import LandingPage from './landingPage'
import LoginPage from './login'
import CreateAccountPage from './createAccount'
import CompaniesPage from './companies'

export default function Pages() {
    return (
        <Fragment>
            <PageContainer>
                <Router primary={false} component={Fragment}>
                    <LandingPage path="/" />
                    <LoginPage path="login" />
                    <CreateAccountPage path="create" />
                    <CompaniesPage path="companies" />
                </Router>
            </PageContainer>
        </Fragment>
    )
}
