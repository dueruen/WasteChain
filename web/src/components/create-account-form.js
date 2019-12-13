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


