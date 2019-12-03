import React, { Fragment } from 'react';

import { LandingComponent } from '../components/landing';

import {  CreateEmployee, GetCompany, ListAllEmployeesInCompany, GetEmployee } from '../api/gql/accounts'
import {  ContinueDoubleSign } from '../api/gql/signature'
import {  GetShipmentDetails, CreateShipment, ListAllShipments } from '../api/gql/shipments'


const LandingPage = () => {
    //CreateShipment("It fucking works, lord have mercy", "2d4036cb-3463-47b1-bae2-69ecf2aec4b1", "here", "2"))
    //GetShipmentDetails("b93f91be-14e1-4644-80bb-ec6754f115b5")
    //CreateEmployee("William M. Buttlicker", "a4b06610-9d71-4851-bbbb-5e91157ef2b1")
    //ContinueDoubleSign("744d3f4c-25c5-4937-8075-0d2d3ffb3d17", "f7bf9284-aa44-449c-9cbc-4f23d904c7f4", "1")


    return (
        <Fragment>
            <LandingComponent/>
        </Fragment>
    )
}

export default LandingPage;
