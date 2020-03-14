export const formOptions = [
  {
    key: 'from_name',
    placeholder: 'Your name',
    inputType: 'text',
    required: true
  },
  {
    key: 'sms_number',
    placeholder: 'Recipients phone number for an SMS',
    inputType: 'tel',
    required: false
  },
  {
    key: 'call_number',
    placeholder: 'Recipients phone number for a phone call',
    inputType: 'tel',
    required: false
  },
  {
    key: 'to_email',
    placeholder: 'Recipients email address',
    inputType: 'email',
    required: false
  },
  {
    key: 'message',
    placeholder: 'What do you want the WUPHF to say?',
    inputType: 'textarea',
    required: true
  },
  {
    key: 'website',
    placeholder: '',
    inputType: 'text',
    required: false
  }
];
