import React, { Fragment } from 'react';
import { Router } from '@reach/router';
import { PageContainer } from '../components';

import LandingPage from './landingPage'
import LoginPage from './login'
import CreateEmployeePage from './createAccount'
import CreateShipmentPage from './shipment/createShipment'

export default function Pages() {
    return (
        <Fragment>
            <PageContainer>
                <Router primary={false} component={Fragment}>
                    <LandingPage path="/" />
                    <LoginPage path="login" />
                    <CreateEmployeePage path="createemployee" />
                    <CreateShipmentPage path="createshipment" />
                </Router>
            </PageContainer>
        </Fragment>
    )
}
