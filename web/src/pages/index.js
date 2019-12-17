import React, { Fragment } from 'react';
import { Router } from '@reach/router';
import { PageContainer } from '../components';

import LandingPage from './landingPage'
import LoginPage from './login'
import CreateEmployeePage from './createAccount'
import CreateCompanyPage from './createCompany'
import CreateShipmentPage from './shipment/createShipment'
import DeatailedShipmentPage from './shipment/shipmentDetails'
import NotFound from '../pages/notFound'

export default function Pages() {
    return (
        <Fragment>
            <PageContainer>
                <Router primary={false} component={Fragment}>
                    <LandingPage path="/" />
                    <LoginPage path="login" />
                    <CreateEmployeePage path="employee/create" />
                    <CreateEmployeePage path="employee/shipments/:employeeID" />
                    <CreateCompanyPage path="company/create" />
                    <CreateShipmentPage path="shipment/create" />
                    <DeatailedShipmentPage path="shipment/:shipmentID"/>
                    <NotFound default/>
                </Router>
            </PageContainer>
        </Fragment>
    )
}
