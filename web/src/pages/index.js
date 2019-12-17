import React, { Fragment } from 'react';
import { Router } from '@reach/router';
import { PageContainer } from '../components';

import LandingPage from './landingPage'
import LoginPage from './login'
import CreateEmployeePage from './createAccount'
import CreateCompanyPage from './createCompany'
import CreateShipmentPage from './shipment/createShipment'
import DeatailedShipmentPage from './shipment/shipmentDetails'

export default function Pages() {
    return (
        <Fragment>
            <PageContainer>
                <Router primary={false} component={Fragment}>
                    <LandingPage path="/" />
                    <LoginPage path="login" />
                    <CreateEmployeePage path="createemployee" />
                    <CreateCompanyPage path="createcompany" />
                    <CreateShipmentPage path="createshipment" />
                    <DeatailedShipmentPage path="shipment/:shipmentID"/>
                </Router>
            </PageContainer>
        </Fragment>
    )
}
