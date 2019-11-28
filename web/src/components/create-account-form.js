import React, { Fragment } from 'react';
import styled from '@emotion/styled'
import useForm from "react-hook-form";

const CreateAccountForm = (props) => {
  const { handleSubmit, register, errors } = useForm();
  const onSubmit = values => {
    console.log(values);
    console.log(props);
    props.create( {variables: {name: values.username}});
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input
        name="email"
        ref={register({
          required: 'Required',
          pattern: {
            value: /^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,4}$/i,
            message: "invalid email address"
          }
        })}
      />
      {errors.email && errors.email.message}

      <input
        name="username"
        ref={register({
          validate: value => value !== "admin" || "Nice try!"
        })}
      />
      {errors.username && errors.username.message}

      <button type="submit">Submit</button>
    </form>
  );
};

export default CreateAccountForm;

// export default function CreateAccountForm(props) {
//     return (
//         <Fragment>
//             <Wrapper>
//                 <LoginBox>
//                     <form action="" method="post" class="form form-login">
//                         <FormField>
//                             <label class="user" for="login-username"><span class="hidden">Username</span></label>
//                             <input id="login-username" type="text" class="form-input" placeholder="Username" required />
//                         </FormField>
//                         <FormField>
//                             <label class="lock" for="login-password"><span class="hidden">Password</span></label>
//                             <input id="login-password" type="password" class="form-input" placeholder="Password" required />
//                         </FormField>
//                         <FormField>
//                             <input type="submit" value="Log in" />
//                         </FormField>
//                     </form>
//                 </LoginBox>
//             </Wrapper>
//         </Fragment>
//     )
// }

// const Wrapper = styled.div`
//     display: flex;
//     flex-direction: column;
// `

// const LoginBox = styled.div`
//     width: '100%';
//     marginBottom: unit * 5;
//     padding: unit * 2.5;
//     position: 'relative';
// `

// const FormField = styled.div`
// display: block;
// margin-bottom: 2rem;
// `
