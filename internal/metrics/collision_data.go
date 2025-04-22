package metrics

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type ParticleData struct {
	Position []float64 `json:"position"`
	Velocity []float64 `json:"velocity"`
	Time     float64   `json:"time"`
}

func (c *ParticleData) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	if err := binary.Write(buf, binary.LittleEndian, uint32(len(c.Position))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, c.Position); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, uint32(len(c.Velocity))); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, c.Velocity); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.LittleEndian, c.Time); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func WriteParticleDataToBinary(filePath string, data []ParticleData) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(file)

	if err := binary.Write(file, binary.LittleEndian, uint32(len(data))); err != nil {
		return err
	}

	for _, entry := range data {
		bin, err := entry.MarshalBinary()
		if err != nil {
			return err
		}
		if err := binary.Write(file, binary.LittleEndian, uint32(len(bin))); err != nil {
			return err
		}
		if _, err := file.Write(bin); err != nil {
			return err
		}
	}

	return nil
}
