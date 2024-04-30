package logg

import "github.com/sirupsen/logrus"

type CustomFormatter struct {
	logrus.TextFormatter
}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+1)
	for k, v := range entry.Data {
		data[k] = v
	}
	data["message"] = entry.Message

	keys := []string{"operation", "host", "port"}
	for _, k := range keys {
		v, ok := data[k]
		if !ok {
			continue
		}
		delete(data, k)
		data[k] = v
	}

	entry.Data = data
	return f.TextFormatter.Format(entry)
}
