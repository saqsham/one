import React, { useState, useMemo, useCallback, useEffect } from 'react';
import PropTypes from 'prop-types';

import { useFormContext } from 'react-hook-form';
import { useMediaQuery } from '@material-ui/core';

import CustomMobileStepper from 'client/components/FormStepper/MobileStepper';
import CustomStepper from 'client/components/FormStepper/Stepper';

const FIRST_STEP = 0;

const FormStepper = ({ steps, initialValue, onSubmit }) => {
  const isMobile = useMediaQuery(theme => theme.breakpoints.only('xs'));
  const { watch, trigger, reset } = useFormContext();

  const [activeStep, setActiveStep] = useState(FIRST_STEP);
  const [formData, setFormData] = useState(initialValue);

  const totalSteps = useMemo(() => steps?.length, [steps]);
  const lastStep = useMemo(() => totalSteps - 1, [totalSteps]);
  const disabledBack = useMemo(() => activeStep === FIRST_STEP, [activeStep]);

  useEffect(() => {
    reset({ ...formData }, { errors: true });
  }, [formData]);

  const handleNext = () => {
    const { id } = steps[activeStep];

    trigger(id).then(isValid => {
      if (!isValid) return;

      if (activeStep === lastStep) {
        onSubmit(formData);
      } else {
        setFormData(prevData => ({ ...prevData, ...watch() }));
        setActiveStep(prevActiveStep => prevActiveStep + 1);
      }
    });
  };

  const handleBack = useCallback(() => {
    if (activeStep <= FIRST_STEP) return;

    setActiveStep(prevActiveStep => prevActiveStep - 1);
  }, [activeStep]);

  return (
    <>
      {/* STEPPER */}
      {isMobile ? (
        <CustomMobileStepper
          totalSteps={totalSteps}
          activeStep={activeStep}
          lastStep={lastStep}
          disabledBack={disabledBack}
          handleNext={handleNext}
          handleBack={handleBack}
        />
      ) : (
        <CustomStepper
          steps={steps}
          activeStep={activeStep}
          lastStep={lastStep}
          disabledBack={disabledBack}
          handleNext={handleNext}
          handleBack={handleBack}
        />
      )}
      {/* FORM CONTENT */}
      {React.useMemo(() => {
        const { id, content: Content } = steps[activeStep];

        return (
          Content && (
            <Content
              step={steps[activeStep]}
              data={formData[id]}
              setFormData={setFormData}
            />
          )
        );
      }, [steps, formData, activeStep, setFormData])}
    </>
  );
};

FormStepper.propTypes = {
  steps: PropTypes.arrayOf(
    PropTypes.shape({
      id: PropTypes.oneOfType([PropTypes.number, PropTypes.string]).isRequired,
      label: PropTypes.string.isRequired,
      content: PropTypes.func.isRequired
    })
  ),
  initialValue: PropTypes.objectOf(PropTypes.any),
  onSubmit: PropTypes.func
};

FormStepper.defaultProps = {
  steps: [],
  initialValue: {},
  onSubmit: console.log
};

export default FormStepper;
